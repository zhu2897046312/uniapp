import request from '@/utils/request';
import type { ApiResponse } from '@/types/api';

/**
 * 图片上传（UniApp 适配版）
 * @param filePath 图片临时路径（通过 uni.chooseImage 获取）
 * @returns 
 */
export const uploadFile = (filePath: string): Promise<string> => {
  return new Promise((resolve, reject) => {
    uni.uploadFile({
      url: 'http://localhost:8080/api/warranty/upload',
      filePath,
      name: 'file',
      success: (uploadRes) => {
        try {
          const res = JSON.parse(uploadRes.data);
          if (res.code === 0) {
            resolve(res.result);
          } else {
            reject(new Error(res.message || '上传失败'));
          }
        } catch (e) {
          reject(new Error('解析响应数据失败'));
        }
      },
      fail: (err) => {
        reject(new Error(err.errMsg));
      }
    });
  });
};

/**
 * 添加质保卡（无需修改）
 */
export const addWarrantyCard = (data: any): Promise<ApiResponse<any>> => {
  return request.post('/warranty/create', data);
};

/**
 * 获取质保卡详细信息（参数优化）
 */
export const getInfoByWarrantyNo = (cardNumber: string): Promise<ApiResponse<any>> => {
  return request.get('/warranty/info', {cardNumber});
};

/**
 * 获取质保卡列表（参数优化）
 */
export const getWarrantyCardList = (data: { cardNumber?: string }): Promise<ApiResponse<any>> => {
  return request.get('/warranty/list', { params: data });
};

/**
 * 机构信息（无需修改）
 */
export const getOraganInfo = (): Promise<ApiResponse<any>> => {
  return request.get('/institution/info');
};

/**
 * 机构登录（无需修改）
 */
export const login = (data: any): Promise<ApiResponse<any>> => {
  return request.post('/login', data);
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
export const forgetPassword = (data: any) : Promise<ApiResponse<any>>=> {
  return request.post('/institution/update-password', data)
}
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