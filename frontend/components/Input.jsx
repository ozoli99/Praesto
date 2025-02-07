import React from "react";

const Input = ({
    label,
    id,
    type = "text",
    value,
    onChange,
    placeholder = "",
    className = "",
}) => {
    return (
        <div className={`flex flex-col ${className}`}>
            {label && (
                <label htmlFor={id} className="mb-1 text-sm font-medium text-gray-700">
                    {label}
                </label>
            )}
            <input
                id={id}
                type={type}
                value={value}
                onChange={onChange}
                placeholder={placeholder}
                className="px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                aria-label={label}
            />
        </div>
    );
};

export default Input;