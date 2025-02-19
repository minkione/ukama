import {
    Grid,
    Accordion,
    Typography,
    AccordionSummary,
    Button,
    AccordionDetails,
} from "@mui/material";
import { colors } from "../../theme";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import QRCode from "qrcode.react";
import { useState } from "react";
interface IESimQR {
    qrCodeId: any;
    description: string;
    isOnBoarding?: boolean;
    title?: string;
    handleClose?: any;
    goToConsole?: any;
}

const ESimQR = ({
    description,
    qrCodeId,
    isOnBoarding = false,
    handleClose,
    goToConsole,
    title,
}: IESimQR) => {
    const [showQrCode, setShowQrCode] = useState(false);

    return (
        <Grid container mb={2}>
            <Grid item xs={12}>
                <Typography variant="h6">{title}</Typography>
            </Grid>
            <Grid item xs={12}>
                <Typography variant="body1">{description}</Typography>
            </Grid>
            <Grid item xs={12}>
                <Accordion
                    sx={{ boxShadow: "none", background: "transparent" }}
                    onChange={(_, isExpanded: boolean) => {
                        setShowQrCode(isExpanded);
                    }}
                >
                    <AccordionSummary
                        expandIcon={<ExpandMoreIcon color="primary" />}
                        sx={{
                            p: 0,
                            m: 0,
                            justifyContent: "flex-start",
                            "& .MuiAccordionSummary-content": {
                                flexGrow: 0.02,
                            },
                        }}
                    >
                        <Typography
                            fontWeight={500}
                            variant="caption"
                            color={colors.primaryMain}
                        >
                            {showQrCode ? "HIDE QR CODE" : "SHOW QR CODE"}
                        </Typography>
                    </AccordionSummary>
                    <AccordionDetails
                        sx={{ p: 0, display: "flex", justifyContent: "center" }}
                    >
                        <QRCode
                            id="qrCodeId"
                            value={qrCodeId}
                            style={{ height: 164, width: 164 }}
                        />
                    </AccordionDetails>
                </Accordion>
            </Grid>

            <Grid item xs={12} container justifyContent="flex-end">
                <Button
                    variant="contained"
                    onClick={isOnBoarding ? goToConsole : handleClose}
                >
                    {isOnBoarding ? "FINISH SETUP" : "CLOSE"}
                </Button>
            </Grid>
        </Grid>
    );
};

export default ESimQR;
