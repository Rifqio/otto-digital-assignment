import { GreetingFormValues } from "../../../schema/greeting-form.schema";

export const GreetingCardPreview: React.FC<{ formValues: Partial<GreetingFormValues> }> = ({ formValues }) => {
    return (
        <div className="w-full aspect-square relative bg-[#FFF5EB] rounded-lg overflow-hidden">
            {formValues.card && (
                <img 
                    src={URL.createObjectURL(formValues.card[0])} 
                    alt="Preview" 
                    className="absolute inset-0 w-full h-full object-cover"
                />
            )}
            
            <div className="absolute inset-0 p-8 flex flex-col">
                <div className="space-y-4">
                    {formValues.dear && (
                        <p className="text-lg font-serif absolute left-52 top-36">
                            {formValues.dear}
                        </p>
                    )}
                    {formValues.message && (
                        <div className="mt-4 whitespace-pre-wrap font-serif text-lg absolute left-32 top-40">
                            {formValues.message}
                        </div>
                    )}
                    {formValues.from && (
                        <div className="mt-auto pt-8">
                            <p className="text-xl font-serif absolute left-48 bottom-44">
                                {formValues.from}
                            </p>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};