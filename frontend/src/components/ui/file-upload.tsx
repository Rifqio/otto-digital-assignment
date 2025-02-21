import {
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Upload } from "lucide-react";
import React, { useState } from "react";
import { FieldValues, UseFormReturn } from "react-hook-form";

interface FileUploadProps<TFormValues extends FieldValues> {
    form: UseFormReturn<TFormValues>;
    name: keyof TFormValues;
    label: string;
}

const FileUpload = <TFormValues extends FieldValues>({
    form,
    name,
    label,
}: FileUploadProps<TFormValues>) => {
    const [isDragging, setIsDragging] = useState<boolean>(false);
    const [thumbnailSrc, setThumbnailSrc] = useState<string>("");

    const handleFileChange = (files: FileList | null): void => {
        const file = files?.item(0);
        if (file) {
            const reader = new FileReader();
            reader.onload = (e: ProgressEvent<FileReader>) => {
                const img = new Image();
                img.src = e.target?.result as string;
                setThumbnailSrc(img.src);
            };
            reader.readAsDataURL(file);
            // @ts-ignore
            form.setValue(name, files, { shouldValidate: true });
        }
    };

    const handleDragOver = (e: React.DragEvent<HTMLDivElement>): void => {
        e.preventDefault();
        setIsDragging(true);
    };

    const handleDragLeave = (e: React.DragEvent<HTMLDivElement>): void => {
        e.preventDefault();
        setIsDragging(false);
    };

    const handleDrop = (e: React.DragEvent<HTMLDivElement>): void => {
        e.preventDefault();
        setIsDragging(false);
        handleFileChange(e.dataTransfer.files);
    };

    return (
        <div className="space-y-4">
            {thumbnailSrc && (
                <div className="flex justify-center">
                    <img
                        src={thumbnailSrc}
                        alt="Preview"
                        className="w-1/2 h-1/2 object-cover rounded-lg"
                    />
                </div>
            )}

            <FormField
                control={form.control}
                name={name}
                // @ts-ignore
                render={({ field: { onChange, value, ...field } }) => (
                    <FormItem>
                        <FormLabel>{label}</FormLabel>
                        <FormControl>
                            <div
                                className={`relative border-2 border-dashed rounded-lg p-8 
                  ${
                      isDragging
                          ? "border-blue-500 bg-blue-50"
                          : "border-gray-300"
                  }
                  hover:border-blue-500 transition-colors duration-200
                  flex flex-col items-center justify-center min-h-[200px]`}
                                onDragOver={handleDragOver}
                                onDragLeave={handleDragLeave}
                                onDrop={handleDrop}
                            >
                                <Input
                                    type="file"
                                    accept="image/*"
                                    onChange={(
                                        e: React.ChangeEvent<HTMLInputElement>
                                    ) => handleFileChange(e.target.files)}
                                    className="absolute inset-0 w-full h-full opacity-0 cursor-pointer"
                                    {...field}
                                />
                                <Upload className="w-10 h-10 text-gray-400 mb-4" />
                                <p className="text-sm font-medium text-gray-700 mb-1">
                                    Browse Files
                                </p>
                                <p className="text-sm text-gray-500">
                                    Drag and drop files here
                                </p>
                            </div>
                        </FormControl>
                        <FormMessage />
                    </FormItem>
                )}
            />
        </div>
    );
};

export default FileUpload;
