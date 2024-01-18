package keeper

import (
	"blog/x/blog/types"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	// Step 1: Get the current post count from the context
    count := k.GetPostCount(ctx)
		// Step 2: Assign the post ID to the post being appended
    post.Id = count
		// Step 3: Create a KVStoreAdapter using the store service and open a KVStore
    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
		// Step 4: Create a new prefixed store (like a dict) with the PostKey prefix
    store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
		// Step 5: Marshal (serialisation) the post object into binary format using the codec (cdc)
    appendedValue := k.cdc.MustMarshal(&post)
		// Step 6: Set the value in the store using the post ID as the key
    store.Set(GetPostIDBytes(post.Id), appendedValue)
		// Step 7: Increment the post count
    k.SetPostCount(ctx, count+1)
		// Step 8: Return the post ID that was assigned to the appended post
    return count
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	// Step 1: Create a KVStoreAdapter using the store service
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Step 2: Create a new prefixed store with an empty prefix
	store := prefix.NewStore(storeAdapter, []byte{})
	// Step 3: Form the byte key using the PostCountKey prefix
	byteKey := types.KeyPrefix(types.PostCountKey)
	// Step 4: Retrieve the raw bytes from the prefixed store using the formed byte key
	bz := store.Get(byteKey)

	// Step 5: Check if the retrieved bytes are nil
	if bz == nil {
			return 0
	}
	 // Step 6: Convert the raw bytes (big-endian encoded uint64) to uint64
	return binary.BigEndian.Uint64(bz)
}

func GetPostIDBytes(id uint64) []byte {
	 // Step 1: Create a byte slice with a length of 8
	bz := make([]byte, 8)
	// Step 2: Convert the uint64 post ID to a big-endian encoded representation and store it in the provided byte slice (bz)
	binary.BigEndian.PutUint64(bz, id)
	// Step 3: Return the resulting byte slice
	return bz
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	// Step 1: Create a KVStoreAdapter using the store service
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	 // Step 2: Create a new prefixed store with an empty prefix
	store := prefix.NewStore(storeAdapter, []byte{})
	// Step 3: Form the byte key using the PostCountKey prefix
	byteKey := types.KeyPrefix(types.PostCountKey)
	// Step 4: Create a byte slice with a length of 8
	bz := make([]byte, 8)
	// Step 5: Convert the uint64 count to big-endian encoded bytes
	binary.BigEndian.PutUint64(bz, count)
	// Step 5: Set the big-endian encoded count in the prefixed store
	store.Set(byteKey, bz)
}

func (k Keeper) GetPost(ctx sdk.Context, id uint64) (val types.Post, found bool) {
	// Step 1: Create a KVStoreAdapter using the store service
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Step 2: Create a new prefixed store with the PostKey prefix
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
	// Step 3: Retrieve the raw bytes from the prefixed store using the post ID key
	b := store.Get(GetPostIDBytes(id))
	// Step 3: Retrieve the raw bytes from the prefixed store using the post ID key
	if b == nil {
			return val, false
	}
	 // Step 5: Unmarshal the raw bytes into a types.Post object
	k.cdc.MustUnmarshal(b, &val)
	// Step 6: Return the unmarshaled types.Post object and set found to true
	return val, true
}

func (k Keeper) SetPost(ctx sdk.Context, post types.Post) {
	// Step 1: Create a KVStoreAdapter using the store service
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Step 2: Create a new prefixed store with the PostKey prefix
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
	// Step 3: Marshal the types.Post object into raw bytes
	b := k.cdc.MustMarshal(&post)
	 // Step 4: Set the marshaled bytes in the prefixed store using the post ID as the key
	store.Set(GetPostIDBytes(post.Id), b)
}

func (k Keeper) RemovePost(ctx sdk.Context, id uint64) {
	// Step 1: Create a KVStoreAdapter using the store service
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Step 2: Create a new prefixed store with the PostKey prefix
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))
	// Step 3: Delete the entry associated with the post ID in the prefixed store
	store.Delete(GetPostIDBytes(id))
}