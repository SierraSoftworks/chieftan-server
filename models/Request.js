"use strict";
exports.RequestSchema = {
    method: /^(GET|POST|PUT|PATCH|DELETE)$/,
    url: String,
    query: { $propertyType: String, $required: false },
    headers: { $propertyType: String, $required: false },
    data: { $required: false, $type: true }
};
