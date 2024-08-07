package gguf_parser

import (
	"slices"

	"github.com/gpustack/gguf-parser-go/util/ptr"
)

type (
	_LLaMACppUsageEstimateOptions struct {
		Architecture        *GGUFArchitectureMetadata
		Tokenizer           *GGUFTokenizerMetadata
		ContextSize         *int32
		InMaxContextSize    bool
		LogicalBatchSize    *int32
		PhysicalBatchSize   *int32
		ParallelSize        *int32
		CacheKeyType        *GGMLType
		CacheValueType      *GGMLType
		OffloadKVCache      *bool
		OffloadLayers       *uint64
		FlashAttention      bool
		MultimodalProjector *LLaMACppUsageEstimate
		Drafter             *LLaMACppUsageEstimate
	}
	LLaMACppUsageEstimateOption func(*_LLaMACppUsageEstimateOptions)
)

// WithArchitecture sets the architecture for the estimate.
//
// Allows reusing the same GGUFArchitectureMetadata for multiple estimates.
func WithArchitecture(arch GGUFArchitectureMetadata) LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		o.Architecture = &arch
	}
}

// WithTokenizer sets the tokenizer for the estimate.
//
// Allows reusing the same GGUFTokenizerMetadata for multiple estimates.
func WithTokenizer(tokenizer GGUFTokenizerMetadata) LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		o.Tokenizer = &tokenizer
	}
}

// WithContextSize sets the context size for the estimate.
func WithContextSize(size int32) LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		if size <= 0 {
			return
		}
		o.ContextSize = &size
	}
}

// WithinMaxContextSize limits the context size to the maximum,
// if the context size is over the maximum.
func WithinMaxContextSize() LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		o.InMaxContextSize = true
	}
}

// WithLogicalBatchSize sets the logical batch size for the estimate.
func WithLogicalBatchSize(size int32) LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		if size <= 0 {
			return
		}
		o.LogicalBatchSize = &size
	}
}

// WithPhysicalBatchSize sets the physical batch size for the estimate.
func WithPhysicalBatchSize(size int32) LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		if size <= 0 {
			return
		}
		o.PhysicalBatchSize = &size
	}
}

// WithParallelSize sets the (decoding sequences) parallel size for the estimate.
func WithParallelSize(size int32) LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		if size <= 0 {
			return
		}
		o.ParallelSize = &size
	}
}

// _GGUFEstimateCacheTypeAllowList is the allow list of cache key and value types.
var _GGUFEstimateCacheTypeAllowList = []GGMLType{
	GGMLTypeF32,
	GGMLTypeF16,
	GGMLTypeQ8_0,
	GGMLTypeQ4_0, GGMLTypeQ4_1,
	GGMLTypeIQ4_NL,
	GGMLTypeQ5_0, GGMLTypeQ5_1,
}

// WithCacheKeyType sets the cache key type for the estimate.
func WithCacheKeyType(t GGMLType) LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		if slices.Contains(_GGUFEstimateCacheTypeAllowList, t) {
			o.CacheKeyType = &t
		}
	}
}

// WithCacheValueType sets the cache value type for the estimate.
func WithCacheValueType(t GGMLType) LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		if slices.Contains(_GGUFEstimateCacheTypeAllowList, t) {
			o.CacheValueType = &t
		}
	}
}

// WithoutOffloadKVCache disables offloading the KV cache.
func WithoutOffloadKVCache() LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		o.OffloadKVCache = ptr.To(false)
	}
}

// WithOffloadLayers sets the number of layers to offload.
func WithOffloadLayers(layers uint64) LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		o.OffloadLayers = &layers
	}
}

// WithFlashAttention sets the flash attention flag.
func WithFlashAttention() LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		o.FlashAttention = true
	}
}

// WithMultimodalProjector sets the multimodal projector estimate usage.
func WithMultimodalProjector(mmp *LLaMACppUsageEstimate) LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		o.MultimodalProjector = mmp
	}
}

// WithDrafter sets the drafter estimate usage.
func WithDrafter(dft *LLaMACppUsageEstimate) LLaMACppUsageEstimateOption {
	return func(o *_LLaMACppUsageEstimateOptions) {
		o.Drafter = dft
	}
}
