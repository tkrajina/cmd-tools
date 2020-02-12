#!/usr/bin/env python3

import subprocess
import datetime
import os
import os.path as path
import re
import shutil

from typing import *


mount_points = "/Volumes"
now_timestamp_str = datetime.datetime.now().strftime('%Y-%m-%dT%H:%M')

def find_clipping_files() -> List[str]:
	files: List[str] = []
	file_name = "documents/My Clippings.txt"
	for dir_name in os.listdir(mount_points):
		fn = path.join(mount_points, dir_name, file_name)
		if path.exists(fn):
			files.append(fn)
	return files

def parse_kindle_notes(file_name: str) -> str:
	keys: List[str] = []
	texts: Dict[str, List[str]] = {}
	unknown = ""

	with open(file_name) as f:
		contents = f.read()
	parts = contents.split("==========")
	for part in parts:
		part = part.strip()
		lines = part.split("\n")
		if len(lines) < 3:
			# unknown += part + "\n"
			pass
		else:
			title = lines[0]
			metadata = lines[1]
			text = "\n".join(lines[2:])
			m = re.match('\-\s+Your\s(.*?)\s+on\s+(.*?)\s*\|.*', metadata)
			if m:
				typ = m.group(1)
				if typ == "Note":
					text = "Note: " + text
				if typ == "Highlight" or typ == "Note":
					typ = "Txt"
				key = f"{title} :: {typ}"
				if not key in texts:
					keys.append(key)
					texts[key] = []
				texts[key].append(text)
			else:
				unknown += part + "\n"

	res = f"\n# Imported {now_timestamp_str}\n"
	keys.sort()
	for key in keys:
		txt = "\n".join(texts[key])
		res += f"\n## {key}\n"
		res += txt + "\n"

	if unknown:
		res += "\n## Unknown\n"
		res += unknown + "\n"

	return res

def main() -> None:
	for file in find_clipping_files():
		contents = parse_kindle_notes(file)
		target_file = "kindle_clippings.md"
		print(contents)
		if input(f"Append to {target_file}? [yn]") == "y":
			with open(target_file, "a") as f:
				f.write(contents)
			if input(f"Clear {file}? [yn]") == "y":
				with open(file, "w") as f:
					f.write("")
				backup_file = f"kindle_clippings_backup_{now_timestamp_str}.txt"
				shutil.copy2(file, backup_file)
				print(f"Backup {file} -> {backup_file}")

main()