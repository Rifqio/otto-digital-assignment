import { forwardRef } from "react";
import { GreetingFormValues } from "../../../schema/greeting-form.schema";

export const GreetingCardPreview = forwardRef<
    HTMLDivElement,
    { formValues: Partial<GreetingFormValues> }
>(({ formValues }, ref) => {
    return (
        <div ref={ref} className="aspect-square relative">
            {formValues.card && (
                <img
                    src={URL.createObjectURL(formValues.card[0])}
                    alt="Preview"
                    className="absolute inset-0 w-full h-full object-cover"
                />
            )}

            <div className="absolute inset-0">
                {formValues.dear && (
                    <div
                        className={`absolute top-[30%]  left-[45%] transform -translate-y-1/2`}
                    >
                        <p className="text-xs md:text-lg sm:text-sm font-serif">
                            {formValues.dear}
                        </p>
                    </div>
                )}

                {formValues.message && (
                    <div className="absolute top-[38%] left-[30%] w-1/2">
                        <p className="md:text-lg sm:text-sm text-xs whitespace-pre-wrap font-serif">
                            {formValues.message}
                        </p>
                    </div>
                )}

                {formValues.from && (
                    <div className="absolute bottom-[41%] left-[41%]">
                        <p className="md:text-lg sm:text-sm font-serif text-xs">
                            {formValues.from}
                        </p>
                    </div>
                )}
            </div>
        </div>
    );
});
