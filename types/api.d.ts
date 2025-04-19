export interface DictResponse {
    customer_name: string;
    institution_name: string;
    surgeon_name: string;
    surgery_date: string;
    prosthesis_model: string;
    code?: number;
    errmsg?: string;
}

export interface ApiResponse<T> {
    code: number;
    data: T;
    message: string;
    errmsg?: string;
} 