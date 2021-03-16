package pipeline

import (
	"errors"

	"dsp-template/api/dbstruct"
	dsp_context "dsp-template/internal/app/dsp/dsp-context"
	dsp_status "dsp-template/internal/app/dsp/dsp-status"
	base_mediafilter "dsp-template/internal/pkg/base-mediafilter"
)

func NewFilteringListPipeline() *FilteringListPipeline {
	return &FilteringListPipeline{
		funcBaseFilter: map[dsp_status.DspStatus]base_mediafilter.MediaFilter{
			dsp_status.DspStatusMediaHardFilterTraffic:     base_mediafilter.MediaHardFilterTraffic(),
			dsp_status.DspStatusMediaHardFilterCountry:     base_mediafilter.MediaHardFilterCountry(),
			dsp_status.DspStatusMediaHardFilterBidCache:    base_mediafilter.MediaHardFilterBidCache(bidCacheTableOption),
			dsp_status.DspStatusMediaHardFilterBannerBType: base_mediafilter.MediaHardFilterBannerBType(),
			dsp_status.DspStatusMediaHardFilterSize:        base_mediafilter.MediaHardFilterSize(),
			dsp_status.DspStatusMediaSoftFilterPackageName: base_mediafilter.MediaSoftFilterPackageName(),
		},
		funcAppFilter: map[dsp_status.DspStatus]base_mediafilter.MediaFilter{
			dsp_status.DspStatusMediaHardFilterUA:         base_mediafilter.MediaHardFilterUA(),
			dsp_status.DspStatusMediaHardFilterDeviceType: base_mediafilter.MediaHardFilterDeviceType(deviceTypeWhiteTableOption),
			dsp_status.DspStatusMediaHardFilterPlatform:   base_mediafilter.MediaHardFilterPlatform(),
			dsp_status.DspStatusMediaHardFilterDeviceID:   base_mediafilter.MediaHardFilterDeviceID(deviceIDWhiteTableOption),
		},
		funcSiteFilter: map[dsp_status.DspStatus]base_mediafilter.MediaFilter{

		},
	}
}

type FilteringListPipeline struct {
	funcAppFilter  map[dsp_status.DspStatus]base_mediafilter.MediaFilter
	funcSiteFilter map[dsp_status.DspStatus]base_mediafilter.MediaFilter
	funcBaseFilter map[dsp_status.DspStatus]base_mediafilter.MediaFilter
}

func (pipe *FilteringListPipeline) Description() string {
	return "FilteringList"
}

func (pipe *FilteringListPipeline) Process(ctx *dsp_context.DspContext) (modelStatus dsp_status.DspStatus, modelErr error) {

	var reserved bool
	reserved = !runMediaHardFilter(ctx, pipe.funcBaseFilter)
	if !reserved {
		return ctx.ModelStatus, errors.New(pipe.Description())
	}

	switch {
	case isFromApp(ctx):
		reserved = !runMediaHardFilter(ctx, pipe.funcAppFilter)
	case isFromSite(ctx):
		reserved = !runMediaHardFilter(ctx, pipe.funcSiteFilter)
	default:
	}
	if !reserved {
		return ctx.ModelStatus, errors.New(pipe.Description())
	}
	return "", nil
}

func isFromApp(ctx *dsp_context.DspContext) bool {
	return ctx.Request.App != nil
}

func isFromSite(ctx *dsp_context.DspContext) bool {
	return ctx.Request.Site != nil
}

func runMediaHardFilter(ctx *dsp_context.DspContext, list map[dsp_status.DspStatus]base_mediafilter.MediaFilter) bool {
	if ctx.Feature == nil {
		writeFilterReason(ctx, dsp_status.DspStatusDefault)
		return true // 过滤流量
	}

	for reason, filter := range list {
		if filter == nil {
			continue
		}
		if filter(ctx.Feature) {
			writeFilterReason(ctx, reason)
			return true // 过滤流量
		}
	}
	return false // 保留流量
}

func writeFilterReason(ctx *dsp_context.DspContext, reason dsp_status.DspStatus) {
	ctx.ModelStatus = reason
}

func deviceTypeWhiteTableOption(feature *dbstruct.Feature) bool {
	// read table

	// if exchange in table; return true
	return true
}

func deviceIDWhiteTableOption(feature *dbstruct.Feature) bool {
	// read table

	// if exchange in table; return true
	return true
}

func bidCacheTableOption(feature *dbstruct.Feature) bool {
	// read table

	// if exchange in table; return true
	return true
}
