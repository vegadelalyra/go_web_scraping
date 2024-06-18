package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	start := time.Now()

	// Create context
	ctx, _ := chromedp.NewContext(context.Background())

	// Set up options to optimize Chromedp initialization and navigation
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("headless", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.9999.999 Safari/537.36"),
	)

	allocCtx, _ := chromedp.NewExecAllocator(ctx, opts...)

	// Create a new chromedp context
	ctx, _ = chromedp.NewContext(allocCtx)

	// Navigate to the page
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://muisca.dian.gov.co/WebRutMuisca/DefConsultaEstadoRUT.faces"),
	); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("~ Navigation took: %s ~\n", time.Since(start))

	start = time.Now()

	// Use the attribute labeldisplay="Nit" to find the corresponding input element
	var inputID string
	if err := chromedp.Run(ctx,
		chromedp.AttributeValue(`input[labeldisplay="Nit"]`, "id", &inputID, nil),
	); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("~ Attribute value lookup took: %s ~\n", time.Since(start))

	start = time.Now()

	// Enter the value into the input field
	value := "65587065"
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`input[id="`+inputID+`"]`, value),
	); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("~ Sending keys took: %s ~\n", time.Since(start))

	start = time.Now()

	// Simulate pressing ENTER (assuming ENTER submits the form)
	if err := chromedp.Run(ctx,
		chromedp.KeyEvent("\r"),
	); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("~ Pressing ENTER took: %s ~\n", time.Since(start))

	start = time.Now()

	// Now retrieve and print the text content from the specified span elements
	fields := map[string]string{
		"primerApellido":  "",
		"segundoApellido": "",
		"primerNombre":    "",
		"otrosNombres":    "",
	}

	for field, value := range fields {
		if err := chromedp.Run(ctx,
			chromedp.Text(`span[id="vistaConsultaEstadoRUT:formConsultaEstadoRUT:`+field+`"]`, &value, chromedp.NodeVisible),
		); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("~ Retrieving %s took: %s ~\n", field, time.Since(start))
	}

	fmt.Println("~ ALL VALUES RETRIEVED AND PRINTED. ~")
}
