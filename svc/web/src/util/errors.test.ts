
import { AxiosError } from 'axios';
import { test, expect } from 'vitest';
import { getAxiosError } from './errors';

interface AxiosErrorCase {
  status: number | undefined;
  messages: Record<number, string> | undefined;
  expected: string;
}

test('getAxiosError returns expected message', () => {
  const cases: AxiosErrorCase[] = [
    {
      status: 400,
      messages: { 400: 'Message' },
      expected: 'Message',
    },
    {
      status: 401,
      messages: { 400: 'Message' },
      expected: 'An unexpected error occurred. Please try again later',
    },
    {
      status: 401,
      messages: {},
      expected: 'An unexpected error occurred. Please try again later',
    },
    {
      status: 401,
      messages: undefined,
      expected: 'An unexpected error occurred. Please try again later',
    },
    {
      status: undefined,
      messages: { 400: 'Message' },
      expected: 'An unexpected error occurred. Please try again later',
    },
  ];

  for (const { status, messages, expected } of cases) {
    const err = { response: { status } } as AxiosError;
    expect(getAxiosError(err, messages)).toBe(expected);
  }
})