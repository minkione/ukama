import {
    NextFn,
    MiddlewareInterface,
} from "type-graphql/dist/interfaces/Middleware";
import { Service } from "typedi";
import { API_METHOD_TYPE } from "../constants";
import { Context, ParsedCookie } from "./types";
import { HTTP401Error, Messages } from "../errors";
import { catchAsyncIOMethod, parseCookie } from ".";
import { ResolverData } from "type-graphql/dist/interfaces/ResolverData";

@Service()
export class Authentication implements MiddlewareInterface<Context> {
    async use({ context }: ResolverData<Context>, next: NextFn): Promise<void> {
        if (context.req.headers.cookie) {
            context.cookie = context.req.headers.cookie;
        } else if (context.req.headers.authorization) {
            context.token = context.req.headers.authorization;
        } else {
            throw new HTTP401Error(Messages.ERR_REQUIRED_HEADER_NOT_FOUND);
        }
        const res = await validateSession(parseCookie(context));
        if (res) return next();
        else throw new HTTP401Error(Messages.ERR_UNAUTHORIZED);
    }
}

const validateSession = async (cookie: ParsedCookie) => {
    const res = await catchAsyncIOMethod({
        type: API_METHOD_TYPE.GET,
        path: "https://auth.dev.ukama.com/.api/sessions/whoami",
        headers: cookie.header,
    });

    return res?.identity?.id ? true : false;
};
