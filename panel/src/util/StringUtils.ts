export function startsWith(str: string, prefix: string): boolean {
  return str.slice(0, prefix.length) === prefix
}

export function endsWith(str: string, suffix: string): boolean {
  if (str.length < suffix.length) {
    return false
  } else {
    return str.slice(str.length - suffix.length) === suffix
  }
}