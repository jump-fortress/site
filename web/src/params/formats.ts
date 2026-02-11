const formats = ['monthly', 'motw', 'archive'];

export function match(value: string) {
  return formats.includes(value);
}
