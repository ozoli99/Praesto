import React from "react";

const Button = ({ children, onClick, className = "", ariaLabel, ...props }) => (
    <button
        onClick={onClick}
        aria-label={ariaLabel}
        className={`
            py-2 px-4
            md:px-6 lg:px-8
            bg-blue-500 hover:bg-blue-600
            text-white rounded
            focus:outline-none focus:ring-2 focus:ring-blue-500
            transition-colors duration-200
            ${className}
        `}
        {...props}
    >
        {children}
    </button>
);

export default Button;
