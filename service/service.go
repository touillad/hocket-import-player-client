package service


type Service interface {
  
	Ingest(ctx context.Context, req)

}