package services

import "context"

var ProductService = productService{}

type productService struct {
}

func (p *productService) mustEmbedUnimplementedProdServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (p *productService) GetProductStock(ctx context.Context, r *ProductRequest) (*ProductResponse, error) {
	stock := p.GetById(r.ProdId)
	return &ProductResponse{ProdStock: stock}, nil
}
func (p *productService) GetById(id int32) int32 {
	return id
}
