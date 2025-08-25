# Decentralized Encrypted Knowledge Sharing Platform Based on Sui Blockchain

## 1. Overall Project Description

This project aims to build a decentralized encrypted knowledge sharing platform based on the Sui blockchain. Creators can upload various types of content on the platform, and users can unlock and access this encrypted content by purchasing AccessTokens. The core concept of the platform is to combine on-chain rights confirmation with off-chain encryption, and to support multiple media content formats. The desktop client will be developed using the Go language and the Fyne framework.

**Core Features:**
*   **On-Chain Rights Confirmation:** Utilize the Sui blockchain to achieve decentralized management of content ownership and access rights.
*   **Off-Chain Encryption:** The content itself is encrypted and stored off-chain to ensure data privacy and security.
*   **Multimedia Content Support:** Supports various content formats such as text, images, videos, and PDFs.
*   **Desktop Client:** Provides a desktop application developed with the Go + Fyne framework, supporting multiple platforms.

## 2. Smart Contract / Move Development

The smart contract will be implemented in the Move language on the Sui blockchain, mainly managing Content and AccessToken objects, and supporting flexible access control and event notifications.

**Main Objects and Functions:**
*   **Content Object:**
    *   Stores the URI (Uniform Resource Identifier) of the content.
    *   Stores the content hash for on-chain verification of content integrity.
    *   Defines the pricing model for the content.
*   **AccessToken Object:**
    *   Records the owner of the AccessToken.
    *   Associates with a Content ID to specify the content it unlocks.
    *   Includes an expiration time to support time-limited access.
    *   Records the revocation status to enable permission withdrawal.
*   **Dynamic Fields:**
    *   Implement multi-level permission management (e.g., basic, premium).
    *   Support composite passes to unlock more advanced or diverse content.
*   **Event Notifications:**
    *   `ContentPublished`: Triggered when content is published.
    *   `TokenIssued`: Triggered when an AccessToken is issued.
    *   `TokenRevoked`: Triggered when an AccessToken is revoked.
*   **On-Chain Hash Verification:** Verifies the integrity and tamper-proof nature of the content by comparing it with the hash stored on-chain.
*   **Contract Support Functions:**
    *   **Subscription:** Supports AccessTokens with monthly or automatic renewal.
    *   **Resale:** Allows users to resell their AccessTokens.
    *   **Secondary Market:** Provides a platform for the resale of AccessTokens and supports creator royalties.

## 3. Encryption and Security

The platform will use advanced encryption technologies to ensure the security of content and key transmission, and to prevent common network attacks.

**Encryption and Security Mechanisms:**
*   **Content Encryption:** Uses the AES-256-GCM algorithm to encrypt content files, providing authenticated encryption.
*   **Key Encryption:** Uses X25519 / Curve25519 public key encryption to decrypt keys, enabling secure key exchange.
*   **Random Nonce and User Signature:** A random Nonce is generated each time content is unlocked, and the user's identity is verified through a signature to prevent replay attacks.
*   **Zero-Knowledge Login (zkLogin):** Supports anonymous access to enhance user privacy.
*   **Temporary Key Mechanism:** Further prevents replay attacks and improves security.

## 4. Desktop Client (Go + Fyne)

The desktop client will be developed using the Go language and the Fyne framework, providing cross-platform support and integrating core functions such as content browsing, purchasing, unlocking, management, and content uploading.

**Main Functional Modules:**
*   **Content Market Browsing:**
    *   Displays a list of content.
    *   Provides a content detail page, including price information.
    *   Integrates a purchase button for users to easily buy AccessTokens.
*   **Purchasing AccessTokens:**
    *   Calls the Sui JSON-RPC interface to interact with the blockchain.
    *   Integrates wallet signature transaction functionality.
*   **Unlocking Content:**
    *   Requests a random number (Nonce) from the gateway service.
    *   Verifies the user's identity with a signature.
    *   Receives the encrypted AES decryption key from the gateway.
    *   Decrypts and displays the content locally.
*   **Managing AccessTokens:**
    *   Views the list of AccessTokens held by the user.
    *   Supports the transfer of AccessTokens.
    *   Provides a function to revoke AccessTokens.
*   **Creator Content Upload:**
    *   Encrypts the content using AES.
    *   Uploads the encrypted content to Walrus (object storage service).
    *   Mints a Content object on the Sui chain.
*   **Asynchronous Processing:** Asynchronously handles network requests and encryption operations to avoid blocking the GUI and improve user experience.
*   **Local Caching:** Caches unlocked content and supports streaming decryption of large files to optimize access speed.

## 5. Gateway Service

The gateway service acts as a bridge between the client, the blockchain, and content storage. It is responsible for key management and user authentication, but does not store the content itself.

**Core Responsibilities:**
*   **Verify User-Signed Random Nonce:** Ensures the legitimacy of requests.
*   **Check On-Chain AccessToken Validity:** Verifies if the user has a valid AccessToken.
*   **Return Encrypted AES Decryption Key:** Sends the encrypted decryption key to the client.
*   **Does Not Store Content:** The gateway only manages decryption keys and does not touch the content itself, enhancing security.
*   **Scalability and High Availability:** Designed to support concurrent requests from multiple users and to be highly available.

## 6. User Experience Enhancement

The platform will be committed to providing an excellent user experience by enhancing user satisfaction through dynamic permissions, multimedia support, and personalized features.

**User Experience Features:**
*   **Dynamic Permission Levels:** Provides different levels of access, such as basic, premium, and collection editions.
*   **Composite AccessTokens:** Allows users to unlock advanced content by combining different AccessTokens.
*   **Local Multimedia Display:** The client supports the local display of various multimedia content such as text, images, videos, and PDFs.
*   **Access History and Favorites Management:** Allows users to easily review and manage visited or interesting content.
*   **Recommendation System:** Intelligently recommends relevant content based on user purchasing behavior and preferences.
*   **Multilingual Support:** The desktop client supports multiple languages.
*   **Cross-Platform Desktop Client:** Supports major operating systems such as Windows, macOS, and Linux.

## 7. Business and Value-Added

The platform will build a sustainable business model to promote a thriving content ecosystem through subscriptions, a secondary market, and incentive mechanisms.

**Business Model and Value-Added Services:**
*   **Subscription-Based AccessTokens:** Offers subscription models with monthly or automatic renewal.
*   **Secondary Market for Reselling AccessTokens:** Allows users to resell AccessTokens on a secondary market and supports creator royalties.
*   **Creator Incentives and Reputation Points:** Incentivizes the creation of high-quality content and establishes a creator reputation system.
*   **Enterprise/Team Edition Knowledge Base Access:** Provides customized knowledge base access solutions for enterprises and teams.
*   **Token Incentive Mechanism and Community Economic Model:** Explores token-based incentive mechanisms to build an active community economy.
## 8. Getting Started

### Prerequisites

- **Go**: Make sure you have Go 1.18 or later installed.
- **Fyne Dependencies**: Follow the [Fyne installation guide](https://developer.fyne.io/started/) to set up the necessary graphics drivers and tools for your operating system.
- **Sui Wallet**: You will need a Sui wallet to interact with the dApp.

### 1. Clone the Repository

```bash
git clone https://github.com/pwh-pwh/AccessSui
cd AccessSui
```

### 2. Configure Your Private Key

You need a Sui account to sign transactions.

1.  Open the `myconfig.go` file.
2.  Replace the placeholder `"your private key"` with your actual Sui account's base64 encoded private key.

    ```go
    // myconfig.go
    package main

    const (
    	PriKey = "YOUR_SUI_PRIVATE_KEY_HERE" // e.g., "privkey:..."
    )
    ```

    **Security Note**: Be extremely careful with your private key. For a real application, use a more secure method for key management, such as environment variables or a dedicated secrets manager.

### 3. Run the Application

To run the application on your desktop:

```bash
go run .
```

### 4. Build for Different Platforms

You can use the `fyne` command-line tool to package your application for different platforms.

- **Desktop**:
  ```bash
  fyne package
  ```
- **Web (WebAssembly)**:
  ```bash
  fyne package -os wasm
  ```
- **Mobile (Android)**:
  ```bash
  fyne package -os android
  ```