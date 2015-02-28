package xml_creator

import (
	"bytes"
	"fmt"
	"strconv"
)

type XML_Message struct {
	msg bytes.Buffer
}

/*** header***/
func (xml *XML_Message) Header() *XML_Message {
	xml.msg.WriteString("<?xml version=\"1.0\"?>\n<commands>\n")
	return xml
}

/*** footer***/
func (xml *XML_Message) Footer() *XML_Message {
	xml.msg.WriteString("</commands>\n")
	return xml
}

/*** disable automation ***/
func (xml *XML_Message) DisableAutomation() *XML_Message {
	xml.msg.WriteString("  <disableAutomation/>\n")
	return xml
}

/*** enable automation ***/
func (xml *XML_Message) EnableAutomation(attributes string) *XML_Message {
	xml.msg.WriteString("  <enableAutomation")
	xml.msg.WriteString(attributes)
	xml.msg.WriteString("/>\n")
	return xml
}

/*** enable automation attributes ***/
func Attr_ea_autoFormat() string          { return " autoFormat=\"true\"" }
func Attr_ea_timestamp() string           { return " timestamp=\"true\"" }
func Attr_ea_cpus(data string) string     { return " cpus=\"" + data + "\"" }     // iqmax cpus => 0 = ptu, 1-6 = bu, 7-9 = su
func Attr_ea_appTypes(data string) string { return " appTypes=\"" + data + "\"" } // iqmax appTypes => callhistory, callmgr, directory, laucherapp, progmgr, speakers
func Attr_ea_items(data string) string    { return " items=\"" + data + "\"" }
func Attr_ea_leds(data string) string     { return " leds=\"" + data + "\"" }

/*** force leds ***/
func (xml *XML_Message) ForceLeds() *XML_Message {
	xml.msg.WriteString("  <forceLeds/>\n")
	return xml
}

/*** force paint ***/
func (xml *XML_Message) ForcePaint() *XML_Message {
	xml.msg.WriteString("  <forcePaint/>\n")
	return xml
}

/*** get trader button status ***/
func (xml *XML_Message) GetTraderButtonStatus(number int) *XML_Message {
	xml.msg.WriteString("  <getTraderButtonStatus number=\"" + strconv.Itoa(number) + "\"/>\n")
	return xml
}

/*** pause ***/
func (xml *XML_Message) Pause(duration int) *XML_Message {
	xml.msg.WriteString("  <pause duration=\"" + strconv.Itoa(duration) + "\"/>\n")
	return xml
}

/*** press a button for a period of time ***/
func (xml *XML_Message) PressButton(cpuid int, button int, duration int) *XML_Message {
	xml.msg.WriteString("  <pressButton" +
		" cpuid=\"" + strconv.Itoa(cpuid) +
		"\" button=\"" + strconv.Itoa(button) +
		"\" duration=\"" + strconv.Itoa(duration) +
		"\"/>\n")
	return xml
}

/*** press buttons for a period of time ***/
func (xml *XML_Message) PressButtons(cpuid int, buttons string, duration int) *XML_Message {
	xml.msg.WriteString("  <pressButton" +
		" cpuid=\"" + strconv.Itoa(cpuid) +
		"\" buttons=\"" + buttons +
		"\" duration=\"" + strconv.Itoa(duration) +
		"\"/>\n")
	return xml
}

/*** return queued messages ***/
func (xml *XML_Message) ReturnQueuedMessages() *XML_Message {
	xml.msg.WriteString("  <returnQueuedMessages/>\n")
	return xml
}

/*** take screenshot ***/
func (xml *XML_Message) TakeScreenshot() *XML_Message {
	xml.msg.WriteString("  <takeScreenshot/>\n")
	return xml
}

/*** print debug info to stdout ***/
func (xml *XML_Message) Debug() *XML_Message {
	fmt.Println(xml.msg.String())
	return xml
}

/*** convert the message to a string ***/
func (xml *XML_Message) String() string {
	return xml.msg.String()
}
