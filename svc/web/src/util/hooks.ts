import { useCallback } from "react";

type ChangeEvent = () => void;

export const useChangePropagator = (handler: (value: string) => void, ...onChange: ChangeEvent[]) => {
  return useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    handler(e.target.value)
    onChange.forEach((fn) => fn());
  }, [handler, onChange]);
};

export const useEnterHandler = (handler: () => void) => {
  return useCallback((e: React.KeyboardEvent) => {
    if (e.key === 'Enter') {
      e.stopPropagation();
      handler();
    }
  }, [handler]);
};