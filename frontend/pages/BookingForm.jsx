import React, { useState } from "react";
import Input from "../components/Input";
import Calendar from "../components/Calendar";
import Button from "../components/Button";

const BookingForm = () => {
    const [date, setDate] = useState("");
    const [name, setName] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        console.log("Booking for", name, "on", date);
    };
    
    return (
        <form onSubmit={handleSubmit} className="p-4 max-w-md mx-auto">
            <Input
                label="Your Name"
                id="name"
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="Enter your name"
                className="mb-4"
            />
            <div className="mt-4">
                <label htmlFor="date" className="block text-sm font-medium text-gray-700">
                    Appointment Date
                </label>
                <Calendar
                    selectedDate={date}
                    onChange={(e) => setDate(e.target.value)}
                    className="mt-1 w-full"
                />
            </div>
            <div className="mt-6">
                <Button type="submit" ariaLabel="Book Appointment">
                    Book Appointment
                </Button>
            </div>
        </form>
    );
};

export default BookingForm;