/* eslint-disable @typescript-eslint/no-unused-vars */
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Button } from "../../../components/ui/button";
import FileUpload from "../../../components/ui/file-upload";
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "../../../components/ui/form";
import { Input } from "../../../components/ui/input";
import { Textarea } from "../../../components/ui/textarea";
import { greetingFormSchema, GreetingFormValues } from "../../../schema/greeting-form.schema";

export function GreetingCardForm() {
    const form = useForm<GreetingFormValues>({
        resolver: zodResolver(greetingFormSchema),
        defaultValues: {
            dear: "",
            message: "",
            from: "",
        },
    });

    function onSubmit(values: z.infer<typeof greetingFormSchema>) {
        console.log(values);
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                <FileUpload<GreetingFormValues> form={form} name="card" label="Card" />
                <FormField
                    control={form.control}
                    name="dear"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Dear</FormLabel>
                            <FormControl>
                                <Input {...field} />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <FormField
                    control={form.control}
                    name="message"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Message</FormLabel>
                            <FormControl>
                                <Textarea {...field} />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <FormField
                    control={form.control}
                    name="from"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>From</FormLabel>
                            <FormControl>
                                <Input {...field} />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <hr />
                <div className="flex justify-center">
                    <Button type="submit" className="bg-green-500">
                        Download
                    </Button>
                </div>
            </form>
        </Form>
    );
}
