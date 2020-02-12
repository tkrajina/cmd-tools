#!/usr/bin/env python3

import tempfile
import datetime
import subprocess
import os.path as path
import os
import time
import shutil
import math

from typing import *

# Screencapture until new...

tmp_dir = tempfile.mkdtemp()
if input("Take screenshots? [yn] ") == "y":
    print("Select window in 5s")
    time.sleep(5)
    running = True
    while running:
        now = datetime.datetime.now()
        screenshot_name = path.join(tmp_dir, now.isoformat() + ".png")
        subprocess.check_output(['screencapture', '-o', '-w', screenshot_name])
        print(f"Screenshot: {screenshot_name}")
        if path.exists(screenshot_name):
            time.sleep(1)
        else:
            running = False

# Copy to directory
directory = ""
while not directory:
    directory = input("Directory: ")
os.makedirs(directory, exist_ok=True)
for file in os.listdir(tmp_dir):
    frm = path.join(tmp_dir, file)
    to = path.join(directory, file)
    print(f"moving {file}: {frm} -> {to}")
    shutil.copy2(frm, to)

if input(f"Open {directory}? [yn] ") == "y":
    subprocess.check_output(["open", directory])

# Convert pages and add page numbers
n=120
w=10
s=100
e=10
prefix = "work_"
if input(f"Convert {directory} to pdf? [yn] ") == "y":
    page_from = 0
    try: page_from = int(input("From page: "))
    except: pass
    page_to = 0
    try: page_to = int(input("To page: "))
    except: pass
    pages = 0
    try: pages = int(input("Total pages: "))
    except: pass

    page_files = []
    for file in os.listdir(directory):
        if not file.endswith(".png"):
            print(f"Ignoring {file}")
            continue
        if file.startswith(prefix):
            print(f"Ignoring {file}")
            continue
        print(f"cropping {file}")
        subprocess.check_output(['convert',
            '-gravity', 'north', '-chop', f'0x{n}',
            '-gravity', 'west', '-chop', f'{w}x0',
            '-gravity', 'south', '-chop', f'0x{s}',
            '-gravity', 'east', '-chop', f'{e}x0',
            '-colorspace', 'Gray', '-quality', '85%',
            path.join(directory, file),
            path.join(directory, prefix + file)
        ])
        page_files.append(path.join(directory, prefix + file))
    page_files.sort()
    if page_from or page_to:
        for n, file in enumerate(page_files):
            page = math.floor(page_from + (n / len(page_files)) * (page_to - page_from))
            print(f"page {page}: {file}")
            label = f"{page}"
            if pages:
                label += f" of {pages}"
            subprocess.check_output([
                'montage',
                '-font',
                "/Library/Fonts/Arial Unicode.ttf",
                '-pointsize',
                '40',
                '-label',
                label,
                '-geometry',
                '+0+0',
                file,
                file,
            ])
    
    outfile = f"{path.basename(os.getcwd())}-{directory}.pdf"
    print(f"Converting {directory} to {outfile}")
    cmd = ["convert"]
    for f in page_files:
        cmd.append(f)
    cmd.append(outfile)
    subprocess.check_output(cmd)

    for file in page_files:
        print(f"Removing {file}")
        os.remove(file)

    if input(f"Open {outfile}? [yn] ") == "y":
        subprocess.check_call(["open", outfile])