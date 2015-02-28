#include <sstream>
#define _FIREWALL_CPP_
#include "Firewall.h"

void openFirewallPort(int m_port)
{
#ifdef PQ2
    std::ostringstream command;
    command << "/sbin/iptables -I tcpBaseIn -p tcp -m tcp --dport " << m_port << " -m state --state NEW,ESTABLISHED -j ACCEPT";
    system (command.str().c_str());
#endif
}

void closeFirewallPort(int m_port)
{
#ifdef PQ2
    std::ostringstream command;
    command << "/sbin/iptables -D tcpBaseIn -p tcp -m tcp --dport " << m_port << " -m state --state NEW,ESTABLISHED -j ACCEPT";
    system (command.str().c_str());
#endif
}

