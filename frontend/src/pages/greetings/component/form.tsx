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
import {
    greetingFormSchema,
    GreetingFormValues,
} from "../../../schema/greeting-form.schema";
import { GreetingCardPreview } from "./preview";

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

    const watch = form.watch();

    return (
        <div className="max-w-2xl mx-auto p-6">
            <div className="space-y-8">
                <div className="rounded-lg overflow-hidden shadow-lg mb-8">
                    {watch.card && (
                        <GreetingCardPreview formValues={watch} />
                    )}
                </div>

                <Form {...form}>
                    <form
                        onSubmit={form.handleSubmit(onSubmit)}
                        className="space-y-6"
                    >
                        <FileUpload<GreetingFormValues>
                            form={form}
                            name="card"
                            label="Card"
                        />

                        <div className="space-y-4">
                            <FormField
                                control={form.control}
                                name="dear"
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Dear</FormLabel>
                                        <FormControl>
                                            <Input {...field} maxLength={50} />
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
                                            <Textarea
                                                {...field}
                                                maxLength={50}
                                                className="min-h-[100px]"
                                            />
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
                                            <Input {...field} maxLength={50} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                        </div>

                        <hr className="my-6" />

                        <div className="flex justify-center">
                            <Button type="submit" className="bg-green-500">
                                Download
                            </Button>
                        </div>
                    </form>
                </Form>
            </div>
        </div>
    );
}
