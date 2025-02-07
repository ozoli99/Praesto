import React from "react";

const Notification = ({
    message,
    type = "info",
    onClose,
    className = "",
}) => {
    let bgColor;
    switch (type) {
        case "success":
            bgColor = "bg-green-500";
            break;
        case "warning":
            bgColor = "bg-yellow-500";
            break;
        case "error":
            bgColor = "bg-red-500";
            break;
        default:
            bgColor = "bg-blue-500";
    }

    return (
        <div
            className={`p-4 rounded text-white ${bgColor} ${className}`}
            role="alert"
            aria-live="assertive"
        >
            <div className="flex justify-between items-center">
                <span>{message}</span>
                {onClose && (
                    <button
                        onClick={onClose}
                        aria-label="Close notification"
                        className="ml-4 text-2xl leading-none"
                    >
                        &times;
                    </button>
                )}
            </div>
        </div>
    );
};

export default Notification;