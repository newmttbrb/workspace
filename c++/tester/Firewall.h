#ifdef _FIREWALL_CPP_

void openFirewallPort(int m_port);
void closeFirewallPort(int m_port);

#else

extern void openFirewallPort(int m_port);
extern void closeFirewallPort(int m_port);

#endif
