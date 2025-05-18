package order

import (
	"context"

	commonpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/common/v1"
	orderpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/order/v1"
	userpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/user/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	orderpbv1.UnimplementedOrderServiceServer
	svc     Service
	userSvc userpbv1.UserServiceClient
}

func NewServer(svc Service, userSvc userpbv1.UserServiceClient) *Server {
	return &Server{svc: svc, userSvc: userSvc}
}

func (s *Server) CreateOrder(ctx context.Context, req *orderpbv1.CreateOrderRequest) (*orderpbv1.CreateOrderResponse, error) {
	client_id, err := uuid.Parse(req.ClientId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID клиента")
	}

	id, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID категории")
	}

	order, err := s.svc.Create(ctx,
		req.Title, req.Description, req.Address,
		req.Longitude, req.Latitude, req.Status,
		req.Price, id, client_id, uuid.Nil,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	clientRes, err := s.userSvc.GetUserById(ctx, &userpbv1.GetUserByIdRequest{UserId: req.ClientId})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var master *commonpbv1.UserData
	if req.MasterId != "" {
		if mres, err := s.userSvc.GetUserById(ctx, &userpbv1.GetUserByIdRequest{UserId: req.MasterId}); err == nil {
			master = mres.User
		}
	}
	return &orderpbv1.CreateOrderResponse{Order: &commonpbv1.OrderData{
		Id:          order.ID.String(),
		Title:       order.Title,
		Description: order.Description,
		Address:     order.Address,
		Longitude:   order.Longitude,
		Latitude:    order.Latitude,
		Status:      order.Status.String(),
		Price:       order.Price,
		CategoryId:  order.CategoryID.String(),
		Client:      clientRes.User,
		Master:      master,
		CreatedAt:   order.CreatedAt.String(),
		UpdatedAt:   order.UpdatedAt.String(),
	}}, nil
}

func (s *Server) GetOrders(ctx context.Context, req *orderpbv1.GetOrdersRequest) (*orderpbv1.GetOrdersResponse, error) {
	client_id, err := uuid.Parse(req.ClientId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID")
	}

	master_id, err := uuid.Parse(req.MasterId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID")
	}

	var categories_ids []uuid.UUID
	if len(req.CategoriesIds) > 0 {
		for _, id := range req.CategoriesIds {
			cid, err := uuid.Parse(id)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID категории")
			}
			categories_ids = append(categories_ids, cid)
		}
	} else {
		categories_ids = nil
	}

	ents, err := s.svc.GetAll(ctx, categories_ids, req.Status, client_id, master_id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	out := make([]*commonpbv1.OrderData, len(ents))
	for i, o := range ents {
		out[i] = &commonpbv1.OrderData{
			Id:          o.ID.String(),
			Title:       o.Title,
			Description: o.Description,
			Address:     o.Address,
			Longitude:   o.Longitude,
			Latitude:    o.Latitude,
			Status:      o.Status.String(),
			Price:       o.Price,
			CategoryId:  o.CategoryID.String(),
			CreatedAt:   o.CreatedAt.String(),
			UpdatedAt:   o.UpdatedAt.String(),
		}
	}
	return &orderpbv1.GetOrdersResponse{Orders: out}, nil
}

func (s *Server) GetOrderById(ctx context.Context, req *orderpbv1.GetOrderByIdRequest) (*orderpbv1.GetOrderByIdResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid UUID")
	}
	o, err := s.svc.Get(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &orderpbv1.GetOrderByIdResponse{Order: &commonpbv1.OrderData{
		Id:          o.ID.String(),
		Title:       o.Title,
		Description: o.Description,
		Address:     o.Address,
		Longitude:   o.Longitude,
		Latitude:    o.Latitude,
		Status:      o.Status.String(),
		Price:       o.Price,
		Client:      &commonpbv1.UserData{Id: o.ClientID.String()},
		CategoryId:  o.CategoryID.String(),
		CreatedAt:   o.CreatedAt.String(),
		UpdatedAt:   o.UpdatedAt.String(),
	}}, nil
}

func (s *Server) UpdateOrder(ctx context.Context, req *orderpbv1.UpdateOrderRequest) (*orderpbv1.GetOrderByIdResponse, error) {
	client_id := uuid.Nil
	master_id := uuid.Nil
	category_id := uuid.Nil

	if req.ClientId != "" {
		client_id_p, err := uuid.Parse(req.ClientId)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID клиента")
		}
		client_id = client_id_p
	}

	if req.MasterId != "" {
		master_id_p, err := uuid.Parse(req.ClientId)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID исполнителя")
		}
		master_id = master_id_p
	}

	if req.CategoryId != "" {
		category_id_p, err := uuid.Parse(req.CategoryId)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID категории")
		}
		category_id = category_id_p
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid UUID")
	}
	ord, err := s.svc.Update(ctx, id,
		req.Title, req.Description, req.Address,
		req.Longitude, req.Latitude, req.Status,
		req.Price, category_id, client_id, master_id,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &orderpbv1.GetOrderByIdResponse{Order: &commonpbv1.OrderData{
		Id:          ord.ID.String(),
		Title:       ord.Title,
		Description: ord.Description,
		Address:     ord.Address,
		Longitude:   ord.Longitude,
		Latitude:    ord.Latitude,
		Status:      ord.Status.String(),
		Price:       ord.Price,
		Client:      &commonpbv1.UserData{Id: ord.ClientID.String()},
		CategoryId:  ord.CategoryID.String(),
		CreatedAt:   ord.CreatedAt.String(),
		UpdatedAt:   ord.UpdatedAt.String(),
	}}, nil
}

func (s *Server) DeleteOrder(ctx context.Context, req *orderpbv1.DeleteOrderRequest) (*orderpbv1.DeleteOrderResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid UUID")
	}
	if err := s.svc.Delete(ctx, id); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &orderpbv1.DeleteOrderResponse{}, nil
}

func (s *Server) GetMyOrders(ctx context.Context, req *orderpbv1.GetMyOrdersRequest) (*orderpbv1.GetMyOrdersResponse, error) {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid UUID")
	}
	ents, err := s.svc.GetAll(ctx, nil, req.Status, id, uuid.Nil)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	out := make([]*commonpbv1.OrderData, len(ents))
	for i, o := range ents {
		out[i] = &commonpbv1.OrderData{
			Id:          o.ID.String(),
			Title:       o.Title,
			Description: o.Description,
			Address:     o.Address,
			Longitude:   o.Longitude,
			Latitude:    o.Latitude,
			Status:      o.Status.String(),
			Price:       o.Price,
			CategoryId:  o.CategoryID.String(),
			CreatedAt:   o.CreatedAt.String(),
			UpdatedAt:   o.UpdatedAt.String(),
		}
	}
	return &orderpbv1.GetMyOrdersResponse{Orders: out}, nil
}

func (s *Server) GetMyFinishedOrders(ctx context.Context, req *orderpbv1.GetMyFinishedOrdersRequest) (*orderpbv1.GetMyFinishedOrdersResponse, error) {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid UUID")
	}
	ents, err := s.svc.GetAll(ctx, nil, "done", id, uuid.Nil)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	out := make([]*commonpbv1.OrderData, len(ents))
	for i, o := range ents {
		out[i] = &commonpbv1.OrderData{
			Id:          o.ID.String(),
			Title:       o.Title,
			Description: o.Description,
			Address:     o.Address,
			Longitude:   o.Longitude,
			Latitude:    o.Latitude,
			Status:      o.Status.String(),
			Price:       o.Price,
			CategoryId:  o.CategoryID.String(),
			CreatedAt:   o.CreatedAt.String(),
			UpdatedAt:   o.UpdatedAt.String(),
		}
	}
	return &orderpbv1.GetMyFinishedOrdersResponse{Orders: out}, nil
}
