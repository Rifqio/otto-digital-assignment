import { zodResolver } from "@hookform/resolvers/zod";
import { render, renderHook, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { useForm } from "react-hook-form";
import { vi } from "vitest";
import GreetingsPage from "../pages/greetings";
import { GreetingCardForm } from "../pages/greetings/component/form";
import {
    greetingFormSchema,
    GreetingFormValues,
} from "../schema/greeting-form.schema";

vi.mock("@hookform/resolvers/zod", () => ({
    zodResolver: () => async (data: any) => ({
        values: data,
        errors: {},
    }),
}));


describe("Greetings Page", () => {
    it("should renders the title from Greetings Page component", () => {
        render(<GreetingsPage />);
        const title = screen.getByText(/Gift Card/i);
        expect(title).toBeInTheDocument();
    });
});

describe("Greeting Card Form", () => {
    it("renders all form fields", () => {
        render(<GreetingCardForm />);

        expect(screen.getByLabelText(/Dear/i)).toBeInTheDocument();
        expect(screen.getByLabelText(/Message/i)).toBeInTheDocument();
        expect(screen.getByLabelText(/From/i)).toBeInTheDocument();
    });

    it("allows user input", async () => {
        renderHook(() =>
            useForm<GreetingFormValues>({
                resolver: zodResolver(greetingFormSchema),
                defaultValues: {
                    dear: "",
                    message: "",
                    from: "",
                },
            })
        );

        render(<GreetingCardForm />);

        const dearInput = screen.getByLabelText(/Dear/i);
        const messageInput = screen.getByLabelText(/Message/i);
        const fromInput = screen.getByLabelText(/From/i);

        await userEvent.type(dearInput, "John");
        await userEvent.type(messageInput, "Happy Birthday!");
        await userEvent.type(fromInput, "Alice");

        expect(dearInput).toHaveValue("John");
        expect(messageInput).toHaveValue("Happy Birthday!");
        expect(fromInput).toHaveValue("Alice");
    });

    it("submits the form successfully", async () => {
        const onSubmit = vi.fn();

        const { result } = renderHook(() =>
            useForm<GreetingFormValues>({
                resolver: zodResolver(greetingFormSchema),
                defaultValues: {
                    dear: "John",
                    message: "Happy Birthday!",
                    from: "Alice",
                },
            })
        );

        const handleSubmit = result.current.handleSubmit(onSubmit);

        render(
            <form onSubmit={handleSubmit}>
                <button type="submit">Download</button>
            </form>
        );

        const button = screen.getByRole("button", { name: /Download/i });
        await userEvent.click(button);

        expect(onSubmit).toHaveBeenCalledWith(
            {
                dear: "John",
                message: "Happy Birthday!",
                from: "Alice",
            },
            expect.anything()
        );
    });
});
