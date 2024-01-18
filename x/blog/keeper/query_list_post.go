package keeper

import (
	"context"

	"blog/x/blog/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListPost(ctx context.Context, req *types.QueryListPostRequest) (*types.QueryListPostResponse, error) {
	// Check if the request is valid
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }
		 // Step 1: Create a KVStoreAdapter using the store service
    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
		// Step 2: Create a new prefixed store with the PostKey prefix
    store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))

		// Step 3: Initialize a slice to hold the list of posts
    var posts []types.Post

		// Step 4: Paginate through the store using the provided pagination parameters
    pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
      // Step 5: Unmarshal the raw bytes into a types.Post object  
				var post types.Post
				// If there is an error during unmarshaling, return the error
        if err := k.cdc.Unmarshal(value, &post); err != nil {
            return err
        }
				 // Step 6: Append the post to the list
        posts = append(posts, post)
				//indicates that there is no error by returning nil. 
        return nil
    })

    if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }

    return &types.QueryListPostResponse{Post: posts, Pagination: pageRes}, nil
}
