import { createAssert } from 'typia';
export interface Payload {
    title: string;
    subtitle: string;
    icon: string;
    background: string;
    titleFont?: string;
    subtitleFont?: string;
}
export const assertPayload = (input: any, errorFactory?: (p: import("typia").TypeGuardError.IProps) => Error): Payload => {
    const __is = (input: any): input is Payload => {
        const $io0 = (input: any): boolean => "string" === typeof input.title && "string" === typeof input.subtitle && "string" === typeof input.icon && "string" === typeof input.background && (undefined === input.titleFont || "string" === typeof input.titleFont) && (undefined === input.subtitleFont || "string" === typeof input.subtitleFont);
        return "object" === typeof input && null !== input && $io0(input);
    };
    if (false === __is(input))
        ((input: any, _path: string, _exceptionable: boolean = true): input is Payload => {
            const $guard = (createAssert as any).guard;
            const $ao0 = (input: any, _path: string, _exceptionable: boolean = true): boolean => ("string" === typeof input.title || $guard(_exceptionable, {
                path: _path + ".title",
                expected: "string",
                value: input.title
            }, errorFactory)) && ("string" === typeof input.subtitle || $guard(_exceptionable, {
                path: _path + ".subtitle",
                expected: "string",
                value: input.subtitle
            }, errorFactory)) && ("string" === typeof input.icon || $guard(_exceptionable, {
                path: _path + ".icon",
                expected: "string",
                value: input.icon
            }, errorFactory)) && ("string" === typeof input.background || $guard(_exceptionable, {
                path: _path + ".background",
                expected: "string",
                value: input.background
            }, errorFactory)) && (undefined === input.titleFont || "string" === typeof input.titleFont || $guard(_exceptionable, {
                path: _path + ".titleFont",
                expected: "(string | undefined)",
                value: input.titleFont
            }, errorFactory)) && (undefined === input.subtitleFont || "string" === typeof input.subtitleFont || $guard(_exceptionable, {
                path: _path + ".subtitleFont",
                expected: "(string | undefined)",
                value: input.subtitleFont
            }, errorFactory));
            return ("object" === typeof input && null !== input || $guard(true, {
                path: _path + "",
                expected: "Payload",
                value: input
            }, errorFactory)) && $ao0(input, _path + "", true) || $guard(true, {
                path: _path + "",
                expected: "Payload",
                value: input
            }, errorFactory);
        })(input, "$input", true);
    return input;
};
