#ifndef NVML_H
#define NVML_H

#include <nvidia/gdk/nvml.h>

#define EXP(x) exp_##x

extern nvmlReturn_t EXP(nvmlInit)(void);
extern nvmlReturn_t EXP(nvmlShutdown)(void);
extern nvmlReturn_t EXP(nvmlDeviceGetTopologyCommonAncestor)(nvmlDevice_t, nvmlDevice_t, nvmlGpuTopologyLevel_t *);


#endif	//NVML_H

