import { Typography } from "../../components/ui/typography";
import { GreetingCardForm } from "./component/form";

function GreetingsPage() {
    return (
        <div className="bg-[#F0F0FA] min-h-screen flex items-center justify-center">
            <div className="bg-white p-8 rounded-lg shadow-lg w-full md:max-w-xl">
                <Typography variant="h3">Gift Card</Typography>
                <hr className="my-8" />
                <GreetingCardForm />
            </div>
        </div>
    );
}

export default GreetingsPage;
