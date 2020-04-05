package gravatar

type gravatarService struct{}

func (s *gravatarService) Generate(ctx content.Context, in *pb.GravatarRequest) *pb.GravatarResponse {

	return &pb.GravatarResponse{Url: gravatar(in.Email, in.Size)}, nil
}

func main() {

}
