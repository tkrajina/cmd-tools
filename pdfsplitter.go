package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	unipdf "github.com/unidoc/unidoc/pdf"
)

func main() {
	pages := flag.String("pages", "", "Pages from-to")
	crop := flag.String("crop", "", "Pages from-to")
	flag.Parse()

	inFile := flag.Arg(0)
	outFile := flag.Arg(1)

	var opts PdfOpts
	if len(*pages) == 1 {
		printDefaultsAndExit("Invalid pages range")
	} else {
		pages, err := getInts(*pages)
		if err != nil {
			printDefaultsAndExit(err.Error())
		}
		if len(pages) != 2 {
			printDefaultsAndExit("Invalid page range")
		}
		opts.pageFrom = pages[0]
		opts.pageTo = pages[1]
	}

	if len(*crop) > 0 {
		crops, err := getInts(*crop)
		fmt.Printf("crops=%#v\n", crops)
		if err != nil {
			printDefaultsAndExit(err.Error())
		}
		if len(crops) != 2 {
			printDefaultsAndExit("Invalid crop percentages")
		}
		opts.percentageVertical = crops[0]
		opts.percentageHorizontal = crops[1]
	}

	fmt.Printf("opts=%#v\n", opts)
	if err := splitPdf(inFile, outFile, opts); err != nil {
		panic(err.Error())
	}
}

func getInts(str string) ([]int, error) {
	delimiter := "-"
	if !strings.Contains(str, delimiter) {
		delimiter = ","
	}
	parts := strings.Split(str, delimiter)
	res := make([]int, len(parts))
	for n := range parts {
		i, err := strconv.ParseInt(parts[n], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("Cannot parse number: %s", parts[n])
		}
		res[n] = int(i)
	}
	return res, nil
}

func printDefaultsAndExit(msg string) {
	fmt.Fprintf(os.Stderr, msg+"\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type PdfOpts struct {
	pageFrom, pageTo                         int
	percentageVertical, percentageHorizontal int
}

func splitPdf(inputPath string, outputPath string, opts PdfOpts) error {
	pdfWriter := unipdf.NewPdfWriter()

	f, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("Error openning %s: %s", inputPath, err.Error())
	}

	defer f.Close()

	pdfReader, err := unipdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return err
	}

	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return err
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	if numPages < opts.pageTo {
		return fmt.Errorf("numPages (%d) < pageTo (%d)", numPages, opts.pageTo)
	}

	for i := opts.pageFrom; i <= opts.pageTo; i++ {
		pageNum := i

		page, err := pdfReader.GetPageAsPdfPage(pageNum)
		if err != nil {
			return err
		}

		bbox, err := page.GetMediaBox()
		if err != nil {
			return err
		}

		// Zoom in on the page middle, with a scaled width and height.
		width := (*bbox).Urx - (*bbox).Llx
		height := (*bbox).Ury - (*bbox).Lly
		newWidth := width * float64(opts.percentageHorizontal) / 100.0
		newHeight := height * float64(opts.percentageVertical) / 100.0
		(*bbox).Llx += newWidth / 2
		(*bbox).Lly += newHeight / 2
		(*bbox).Urx -= newWidth / 2
		(*bbox).Ury -= newHeight / 2
		page.MediaBox = bbox

		pageObj := page.GetPageAsIndirectObject()

		err = pdfWriter.AddPage(pageObj)
		if err != nil {
			return err
		}
	}

	fWrite, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		return err
	}

	return nil
}
