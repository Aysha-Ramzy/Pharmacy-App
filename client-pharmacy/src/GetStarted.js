import React from 'react';
import { 
    Box,
    Grid,
    styled,
    Typography,
} from '@mui/material';
import Title from './Title';
// img
import pharmacyLandingImage from './resources/pharmacy_landing.png';

const GetStarted = () => {

    const CustomGridItem = styled(Grid) ({
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
    });

    const CustomTypography = styled(Typography) ({
        fontSize: '1.1rem',
        textAlign: 'start',
        lineHeight: '1.5',
        color: '#515151',
        marginTop: '1.5rem',
    });

    return (
        <Grid container spacing={{ xs: 4, sm: 4, md: 0 }}   
            sx={{
                py: 10,
                px: 2,
            }}
        >
            <CustomGridItem item xs={12} sm={8} md={6} 
                component='section'
            >
                <Box component='article'
                    sx={{
                        px: 4,
                    }}
                >
                    <Title
                        text={
                            'We simplify the process for pharmacies'
                        }
                        textAlign={'start'}
                    />
                    <CustomTypography>
                    Keep a tight grip on your inventory, prescriptions, and finances.< br /> E-Pharmacy puts you in control, providing real-time insights and tools to make informed decisions.
                    </CustomTypography> 
                </Box>
            </CustomGridItem>
            
            <Grid item xs={12} sm={4} md={6}>
                <img src={pharmacyLandingImage} alt="" 
                    style={{
                        width: '100%',
                    }}
                />
            </Grid>

        

            <CustomGridItem item xs={12} sm={8} md={6}
                sx={{
                    order: { xs: 3, sm: 3, md: 4 },
                }}
            >
               
            </CustomGridItem>
        </Grid>
    );
};

export default GetStarted;
