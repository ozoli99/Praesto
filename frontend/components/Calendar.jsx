import React from "react";

const Calendar = ({ selectedDate, onChange, className = "" }) => {
    return (
        <input
            type="date"
            value={selectedDate}
            onChange={onChange}
            className={`px-3 py-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500 ${className}`}
            aria-label="Select date"
        />
    );
};

export default Calendar;