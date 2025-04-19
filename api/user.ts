import request from '@/utils/request';
import type { ApiResponse } from '@/types/api';

/**
 * 获取手机验证码（查看质保卡）
 */
export const getFindInfoVerifyCode = (data: any): Promise<ApiResponse<any>> => {
  return request.post('/client/warrantyCard/sendCode', data);
};

/**
 * 提交表单查看质保卡信息
 */
export const getInfoByWarrantyNo = (data: any): Promise<ApiResponse<any>> => {
  return request.post('/client/warrantyCard/info', data);
};

/**
 * 校验初始密码（绑定手机号）
 */
export const bindPhone_password = (data: any): Promise<ApiResponse<any>> => {
  return request.post('/client/warrantyCard/login', data);
};

/**
 * 获取绑定手机号的验证码
 */
export const getBindPhoneVerifyCode = (data: any): Promise<ApiResponse<any>> => {
  return request.post('/client/warrantyCard/bindsendcode', data);
};

/**
 * 绑定手机号
 */
export const bindPhone_code = (data: any): Promise<ApiResponse<any>> => {
  return request.post('/client/warrantyCard/bindmobile', data);
};