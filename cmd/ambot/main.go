package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/ashokhin/am4bot/auth"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Product struct {
	Name, Price, Image, URL string
}

var logger slog.Logger

func init() {
	logger = *slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(&logger)
}

func main() {
	timeStart := time.Now()
	a := auth.New(&logger)
	a.Auth()
	// initialize the Chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var products []Product

	// create a channel to receive products
	productChan := make(chan Product)
	done := make(chan bool)

	// start a goroutine to collect products
	go func() {
		for product := range productChan {
			logger.Info("product found", "name", product.Name, "price", product.Price)
			products = append(products, product)
		}
		done <- true
	}()

	// navigate and scrape
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.scrapingcourse.com/infinite-scrolling"),
		scrapeProducts(productChan),
	)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	close(productChan)
	<-done

	// print results
	logger.Info("scrape completed", "scraped_products", len(products))
	for _, p := range products {
		logger.Info("product info", "name", p.Name, "price", p.Price, "image_url", p.Image, "product_url", p.URL)
	}

	logger.Info("checks success", "elapsed_time", fmt.Sprintf("%v", time.Since(timeStart)))
}

func scrapeProducts(productChan chan<- Product) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		var previousHeight int
		for {
			// get all product nodes
			var nodes []*cdp.Node
			if err := chromedp.Nodes(".product-item", &nodes).Do(ctx); err != nil {
				return err
			}

			// extract data from each product
			for _, node := range nodes {
				var product Product

				// using chromedp's node selection to extract data
				if err := chromedp.Run(ctx,
					chromedp.Text(".product-name", &product.Name, chromedp.ByQuery, chromedp.FromNode(node)),
					chromedp.Text(".product-price", &product.Price, chromedp.ByQuery, chromedp.FromNode(node)),
					chromedp.AttributeValue("img", "src", &product.Image, nil, chromedp.ByQuery, chromedp.FromNode(node)),
					chromedp.AttributeValue("a", "href", &product.URL, nil, chromedp.ByQuery, chromedp.FromNode(node)),
				); err != nil {
					continue
				}

				// clean price text
				product.Price = strings.TrimSpace(product.Price)

				// send product to channel if not empty
				if product.Name != "" {
					productChan <- product
				}
			}

			// scroll to bottom
			var height int
			if err := chromedp.Evaluate(`document.documentElement.scrollHeight`, &height).Do(ctx); err != nil {
				return err
			}

			// break if we've reached the bottom (no height change after scroll)
			if height == previousHeight {
				break
			}
			previousHeight = height

			// scroll and wait for content to load
			if err := chromedp.Run(ctx,
				chromedp.Evaluate(`window.scrollTo(0, document.documentElement.scrollHeight)`, nil),
				chromedp.Sleep(3*time.Second), // Wait for new content to load
			); err != nil {
				return err
			}
		}
		return nil
	}
}
