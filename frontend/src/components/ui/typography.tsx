interface TypographyProps {
    variant: "h1" | "h2" | "h3" | "h4" | "p";
    children: React.ReactNode;
}

export const Typography = ({ variant, children }: TypographyProps) => {
    const Component = variant === "h1" ? "h1" : variant === "h2" ? "h2" : "h3";
    const h1Class =
        "scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl";
    const h2Class =
        "scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight first:mt-0";
    const h3Class = "scroll-m-20 text-2xl font-semibold tracking-tight";
    const h4Class = "scroll-m-20 text-xl font-semibold tracking-tight";
    const pClass = "leading-7 [&:not(:first-child)]:mt-6";

    return (
        <Component
            className={`
            ${variant === "h1" ? h1Class : ""}
            ${variant === "h2" ? h2Class : ""}
            ${variant === "h3" ? h3Class : ""}
            ${variant === "h4" ? h4Class : ""}
            ${variant === "p" ? pClass : ""}
        `}
        >
            {children}
        </Component>
    );
};
