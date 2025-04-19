// 后端响应数据的接口定义
export interface DictResponse {
    customer_name: string;
    institution_name: string;
    surgeon_name: string;
    surgery_date: string;
    prosthesis_model: string;
}

// API 通用响应格式
export interface ApiResponse<T> {
    code: number;
    message: string;
    result: T;
}
