import { z } from 'zod';

export const greetingFormSchema = z.object({
    dear: z.string().nonempty().max(50),
    message: z.string().nonempty().max(200),
    from: z.string().nonempty().max(50),
    card: z.instanceof(FileList)
})

export type GreetingFormValues = z.infer<typeof greetingFormSchema>;
