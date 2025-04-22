import not_login_event from '@/events/not_login_event';

// 定义全局的基础 URL 变量，将这里的 IP 地址和端口修改为你需要的
const baseUrl = 'http://localhost:8080/api'; 

// UniApp 请求封装
const request = {
  // 统一请求方法
  async request(options: UniApp.RequestOptions): Promise<any> {
    // 获取 token（使用 uni.getStorageSync）
    const token = uni.getStorageSync('Authorization');
    
    // 设置请求头
    const header = {
      'Content-Type': 'application/json',
      ...options.header,
    };
    
    if (token) {
      header['Authorization'] = `${token}`;
    }

    // 发起请求，使用新的 baseUrl 拼接地址
    return new Promise((resolve, reject) => {
      uni.request({
        url: baseUrl + options.url, // 拼接 baseURL
        method: options.method || 'GET',
        data: options.data || {},
        header,
        success: (res) => {
          // 处理 401 或自定义 code（如 18000）
          if (res.statusCode === 401 || res.data?.code === 18000) {
            not_login_event.tigger(); // 触发未登录事件
            reject(res.data?.message || '未登录');
          } else {
            resolve(res.data); // 返回数据
          }
        },
        fail: (err) => {
          reject(err);
        },
      });
    });
  },

  // 封装 GET 方法
  get<T>(url: string, params?: any, options?: Partial<UniApp.RequestOptions>): Promise<T> {
    return this.request({
      url,
      method: 'GET',
      data: params,
      ...options,
    });
  },

  // 封装 POST 方法
  post<T>(url: string, data?: any, options?: Partial<UniApp.RequestOptions>): Promise<T> {
    return this.request({
      url,
      method: 'POST',
      data,
      ...options,
    });
  },
};

export default request;