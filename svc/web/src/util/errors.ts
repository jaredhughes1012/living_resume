import { AxiosError } from "axios";

type MessageMap = Record<number, string> | undefined;

export const getAxiosError = (err: unknown, messages: MessageMap): string => {
  const status = (err as AxiosError)?.response?.status;
  const message = (status && messages?.[status]) || "An unexpected error occurred. Please try again later";
  return message;
}