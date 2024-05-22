
import { expect, test, vi } from 'vitest';
import { render, screen } from '@testing-library/react';
import InputBar from './InputBar';
import userEvent from '@testing-library/user-event'
import '@testing-library/jest-dom'

test('Chat input with enter key', async () => {
  const message = 'Hello, world!';
  const user = userEvent.setup();
  const onSubmit = vi.fn();

  render(<InputBar onSubmit={onSubmit} />);

  const input = await screen.findByLabelText('Ask a question');
  await user.type(input, message);
  await user.keyboard('{enter}');

  expect(screen.queryByText(message)).not.toBeInTheDocument();
  expect(onSubmit).toHaveBeenCalledWith({ message });
});

test('Chat input with submit button', async () => {
  const message = 'Hello, world!';
  const user = userEvent.setup();
  const onSubmit = vi.fn();

  render(<InputBar onSubmit={onSubmit} />);

  const input = await screen.findByLabelText('Ask a question');
  await user.type(input, message);

  const button = await screen.findByRole('button');
  await user.click(button);

  expect(screen.queryByText(message)).not.toBeInTheDocument();
  expect(onSubmit).toHaveBeenCalledWith({ message });
});