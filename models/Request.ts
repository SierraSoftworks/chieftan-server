export interface Request {
    method: "GET" | "POST" | "PUT" | "PATCH" | "DELETE";
    url: string;
    query?: {
        [key: string]: string;
    };
    headers?: {
        [header: string]: string;
    };
    data?: string|Object;
}

export const RequestSchema = {
    method: /^(GET|POST|PUT|PATCH|DELETE)$/,
    url: String,
    query: { $propertyType: String, $required: false },
    headers: { $propertyType: String, $required: false },
    data: { $required: false, $type: true }
};