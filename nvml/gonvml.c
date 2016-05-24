#include "gonvml.h"


nvmlReturn_t EXP(nvmlInit)(void){
				return nvmlInit();
}

nvmlReturn_t EXP(nvmlShutdown)(void){
				return nvmlShutdown();
}

nvmlReturn_t EXP(nvmlDeviceGetTopologyCommonAncestor)(nvmlDevice_t dev1, nvmlDevice_t dev2, nvmlGpuTopologyLevel_t *info){
	return nvmlDeviceGetTopologyCommonAncestor(dev1, dev2, info);
}

