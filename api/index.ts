import request from '@/utils/request';
import type { DictResponse, ApiResponse } from '@/types/api';

export default {
  getDictSelect(params: { serialNumber: string }): Promise<ApiResponse<DictResponse>> {
    return request.get('/client/othpcd/dict/select', { params });
  },
};