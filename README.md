# Go API client for Aha! (aha.io)

[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

 [build-status-svg]: https://api.travis-ci.org/grokify/go-aha.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/go-aha
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-aha
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/go-aha
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/go-aha
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-aha/blob/master/LICENSE

Articles that matter on social publishing platform

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.

It also comes with a Postman Collection in the [`postman`](postman) folder.

## Installation

```bash
$ go get github.com/grokify/go-aha/...
```

## Usage

See:

* [`examples/get_features`](examples/get_features)
* [`examples/get_products`](examples/get_products)

## Documentation for API Endpoints

All URIs are relative to *https://secure.aha.io/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FeaturesApi* | [**GetFeature**](docs/FeaturesApi.md#getfeature) | **Get** /features/{feature_id} | 
*FeaturesApi* | [**GetFeatures**](docs/FeaturesApi.md#getfeatures) | **Get** /features | Get all features
*FeaturesApi* | [**GetReleaseFeatures**](docs/FeaturesApi.md#getreleasefeatures) | **Get** /releases/{release_id}/features | Get all features for a release
*FeaturesApi* | [**UpdateFeature**](docs/FeaturesApi.md#updatefeature) | **Put** /features/{feature_id} | Update a feature&#39;s custom fields with tag-like value
*ProductsApi* | [**GetProduct**](docs/ProductsApi.md#getproduct) | **Get** /products/{product_id} | Products API
*ProductsApi* | [**GetProducts**](docs/ProductsApi.md#getproducts) | **Get** /products | Products API
*ReleasesApi* | [**GetRelease**](docs/ReleasesApi.md#getrelease) | **Get** /releases/{release_id} | 
*ReleasesApi* | [**UpdateProductRelease**](docs/ReleasesApi.md#updateproductrelease) | **Put** /products/{product_id}/releases/{release_id} | Update a release


## Documentation For Models

 - [Feature](docs/Feature.md)
 - [FeatureMeta](docs/FeatureMeta.md)
 - [FeatureUpdate](docs/FeatureUpdate.md)
 - [FeatureWrap](docs/FeatureWrap.md)
 - [FeaturesResponse](docs/FeaturesResponse.md)
 - [Pagination](docs/Pagination.md)
 - [Product](docs/Product.md)
 - [ProductMeta](docs/ProductMeta.md)
 - [ProductResponse](docs/ProductResponse.md)
 - [ProductsResponse](docs/ProductsResponse.md)
 - [Release](docs/Release.md)
 - [ReleaseUpdate](docs/ReleaseUpdate.md)
 - [ReleaseUpdateWrap](docs/ReleaseUpdateWrap.md)
 - [ReleaseWrap](docs/ReleaseWrap.md)
 - [ReleasesResponse](docs/ReleasesResponse.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author



