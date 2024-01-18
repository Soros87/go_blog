package keeper

import (
	"context"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowPost(goCtx context.Context, req *types.QueryShowPostRequest) (*types.QueryShowPostResponse, error) {
    //Check if the Request is Valid:
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }
    //Unwrap the Context to Get the Cosmos SDK Context:
    ctx := sdk.UnwrapSDKContext(goCtx)
    //Retrieve the Post from the Keeper Using the Provided Post ID:
    post, found := k.GetPost(ctx, req.Id)
    //Check if the Post Was Found:
    if !found {
        return nil, sdkerrors.ErrKeyNotFound
    }
    //Return a Response Containing the Retrieved Post
    return &types.QueryShowPostResponse{Post: &post}, nil
}