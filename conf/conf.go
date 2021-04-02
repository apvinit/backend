package conf

// DynamicLinkAPIKey is key required for genrating dynamic links
const DynamicLinkAPIKey = "AIzaSyD5nKQMrrZnRN089ynNPQ6z0AJPxe1j-hA"

// AndroidPackageName for the app
const AndroidPackageName = "xyz.codingabc.jobadda.debug"

// AndroidMinPackageVersionCode specifies which app can handle deep link
const AndroidMinPackageVersionCode = "7"

// DomainURIPrefix (It is used for generating shortlink)
// Eg. "https://sjobadda.page.link/sd4T"
const DomainURIPrefix = "https://sjobadda.page.link"

////////////////////////////////////////////////////////////////////////////////////////////////////////////
/// ------------------------------------  Release Config Values ---------------------------------------- ///
////////////////////////////////////////////////////////////////////////////////////////////////////////////

// DeepLinkDomain for deeplink redirect url (same for debug and prod)
const DeepLinkDomain = "https://jbda.in/"

// DefaultImageLink is used when imageLink is not passed in the create post request (same for debug & prod)
const DefaultImageLink = "https://firebasestorage.googleapis.com/v0/b/job-adda.appspot.com/o/images%2Fbanner.jpg?alt=media&token=2081e3ff-b79e-4cd9-9b6b-ae65a74336b8"

const ServerPort = ":8080"

// const DynamicLinkAPIKey = "AIzaSyCAR4-s7t0iN8ZGQjUl__1Q0miAA3x7CuU"

// const AndroidPackageName = "xyz.codingabc.jobadda"

// const AndroidMinPackageVersionCode = "7"

// const DomainURIPrefix = "https://jbda.page.link"
