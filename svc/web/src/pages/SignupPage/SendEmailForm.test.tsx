import { expect, test, vi } from 'vitest';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom'
import SendEmailForm from './SendEmailForm';
import userEvent from '@testing-library/user-event'
import { MemoryRouter } from 'react-router-dom';

test('SendEmailForm submit on button press', async () => {
  const user = userEvent.setup();
  const email = 'test@test.com';
  const onContinue = vi.fn();
  render(
    <MemoryRouter>
      <SendEmailForm onContinue={onContinue} />
    </MemoryRouter>
  );

  const emailInput = screen.getByLabelText('Email');
  await user.type(emailInput, email);
  await user.click(screen.getByText('Continue'));

  expect(onContinue).toHaveBeenCalledWith({ email });
});

test('SendEmailForm submit on enter', async () => {
  const user = userEvent.setup();
  const email = 'test@test.com';
  const onContinue = vi.fn();
  render(
    <MemoryRouter>
      <SendEmailForm onContinue={onContinue} />
    </MemoryRouter>
  );

  const emailInput = screen.getByLabelText('Email');
  await user.type(emailInput, email);
  await user.type(emailInput, '{enter}');

  expect(onContinue).toHaveBeenCalledWith({ email });
});