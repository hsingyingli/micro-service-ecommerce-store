const imageBaseURL = process.env.NEXT_PUBLIC_IMAGE_BASE_URL || ""

const getImageUrl = (imageName: string) => {
  return `${imageBaseURL}/${imageName}`
}

export {
  getImageUrl
}
