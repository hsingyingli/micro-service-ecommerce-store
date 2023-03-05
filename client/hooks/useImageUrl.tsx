
const useImageUrl = (imageName: string): string => {
  const imageBaseURL = process.env.NEXT_PUBLIC_IMAGE_BASE_URL || ""
  return `${imageBaseURL}/${imageName}`
}


export {
  useImageUrl
}
