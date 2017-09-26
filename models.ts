export interface Settings {
    hasInitialized: boolean;
    allowedDomains: string;
    requireModeration: boolean;
    checkUrls: boolean;
    smtpHost: string;
    smtpPass: string;
    smtpPort: number;
    smtpFrom: string;
    smtpSecure: boolean;
    backupRep: string;
    backupToken: string;
}