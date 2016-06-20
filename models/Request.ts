export interface RequestDoc {
    method: "GET" | "POST" | "PUT" | "PATCH" | "DELETE";
    url: string;
    headers?: {
        [header: string]: string;
    };
    data?: string|Object;
}

export const RequestSchema = {
    method: /^(GET|POST|PUT|PATCH|DELETE)$/,
    url: String,
    headers: { $propertyType: String, $required: false },
    data: { $required: false, $type: true }
};