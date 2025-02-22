import { z } from "zod";

export const greetingFormSchema = z.object({
    dear: z
        .string()
        .nonempty({ message: "Recipient must be filled" })
        .max(50, { message: "Recipient must be less than 50 characters" }),
    message: z
        .string()
        .nonempty({ message: "Message must be filled" })
        .max(50, { message: "Message must be less than 200 characters" }),
    from: z.string().nonempty({ message: "From must be filled" }).max(50, { message: "From must be less than 50 characters" }),
    card: z.instanceof(FileList, { message: "Card must be uploaded" }),
});

export type GreetingFormValues = z.infer<typeof greetingFormSchema>;
