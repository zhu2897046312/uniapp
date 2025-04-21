import request from '@/utils/request';
import type { ApiResponse } from '@/types/api';

/**
 * 图片上传（UniApp 适配版）
 * @param filePath 图片临时路径（通过 uni.chooseImage 获取）
 * @returns 
 */
export const uploadFile = (filePath: string): Promise<ApiResponse<any>> => {
	const formData = new FormData();
	formData.append('file', filePath);
  return request.post('/warranty/upload',formData,{
    header: {
      'Content-Type': 'multipart/form-data'
    }
  });
};

/**
 * 添加质保卡（无需修改）
 */
export const addWarrantyCard = (data: any): Promise<ApiResponse<any>> => {
  return request.post('/institution/warrantyCard/create', data);
};

/**
 * 获取质保卡详细信息（参数优化）
 */
export const getInfoByWarrantyNo = (cardNumber: string): Promise<ApiResponse<any>> => {
  return request.get('/institution/warrantyCard/info', { params: { cardNumber } });
};

/**
 * 获取质保卡列表（参数优化）
 */
export const getWarrantyCardList = (data: { cardNumber?: string }): Promise<ApiResponse<any>> => {
  return request.get('/institution/warrantyCard/list', { params: data });
};

/**
 * 机构信息（无需修改）
 */
export const getOraganInfo = (): Promise<ApiResponse<any>> => {
  return request.get('/institution/othpcd/info');
};

/**
 * 机构登录（无需修改）
 */
export const login = (data: any): Promise<ApiResponse<any>> => {
  return request.post('/institution/othpcd/login', data);
};

/**
 * 获取手机验证码（拼写修正）
 */
export const getOrgenVerifyCode = (data: any): Promise<ApiResponse<any>> => {
  return request.post('/institution/othpcd/sendCode', data);
};

/**
 * 修改密码（函数名语义化）
 */
export const changePassword = (data: any): Promise<ApiResponse<any>> => {
  return request.post('/institution/othpcd/change', data);
};

// 导出所有接口
export default {
  uploadFile,
  addWarrantyCard,
  getInfoByWarrantyNo,
  getWarrantyCardList,
  getOraganInfo,
  login,
  getOrgenVerifyCode,
  changePassword
};