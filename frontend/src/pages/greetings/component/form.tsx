import { zodResolver } from "@hookform/resolvers/zod";
import html2canvas from "html2canvas";
import { useRef } from "react";
import { useForm } from "react-hook-form";
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
    const watch = form.watch();
    const previewRef = useRef<HTMLDivElement>(null);

    async function onSubmit() {
        if (!previewRef.current) return;

        await new Promise((resolve) => setTimeout(resolve, 100));

        const canvas = await html2canvas(previewRef.current, {
            backgroundColor: null,
            scale: 2,
            useCORS: true,
            logging: false,
            allowTaint: true,
            scrollX: 0,
            scrollY: 0,
        });

        const image = canvas.toDataURL("image/png", 1.0);

        const link = document.createElement("a");
        link.href = image;
        link.download = "greeting-card.png";
        link.click();
    }
    return (
        <div className="max-w-2xl mx-auto p-6">
            <div className="space-y-8">
                {watch.card && (
                    <GreetingCardPreview ref={previewRef} formValues={watch} />
                )}

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
                                            <Input {...field} maxLength={20} />
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
                                                maxLength={45}
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
                                            <Input {...field} maxLength={20} />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                        </div>

                        <hr className="my-6" />

                        <div className="flex justify-center">
                            <Button
                                type="submit"
                                className="bg-green-500 hover:cursor-pointer hover:bg-green-800"
                            >
                                Download
                            </Button>
                        </div>
                    </form>
                </Form>
            </div>
        </div>
    );
}
